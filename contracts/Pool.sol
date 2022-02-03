// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./LPToken.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/utils/math/Math.sol";

contract Pool is LPToken {
    IERC20 private _tomToken;
    IERC20 private _jerryToken;

    constructor(address tomToken, address jerryToken) LPToken() {
        _tomToken = IERC20(tomToken);
        _jerryToken = IERC20(jerryToken);
    }

    function addLiquidity(uint256 tomAmount, uint256 jerryAmount) public {
        address poolAddr = address(this);
        uint256 reserveTom = _tomToken.balanceOf(poolAddr);
        uint256 reserveJerry = _jerryToken.balanceOf(poolAddr);

        _tomToken.transferFrom(msg.sender, poolAddr, tomAmount);
        _jerryToken.transferFrom(msg.sender, poolAddr, jerryAmount);

        uint256 newReserveTom = _tomToken.balanceOf(poolAddr);
        uint256 newReserveJerry = _jerryToken.balanceOf(poolAddr);

        /*
            reserveTom * reserveJerry = k
            newReserveTom * newReserveJerry >= k
        */
        require(
            newReserveTom * newReserveJerry >= reserveTom * reserveJerry,
            "invalid k"
        );

        uint256 totalSupply = totalSupply();
        uint256 lpToken = 0;
        if (totalSupply == 0) {
            lpToken = tomAmount;
        } else {
            lpToken = Math.min(
                (tomAmount * totalSupply) / reserveTom,
                (jerryAmount * totalSupply) / reserveJerry
            );
        }

        require(lpToken > 0, "lpToken should not be 0");
        _mint(msg.sender, lpToken);
    }

    function removeLiquidity(uint256 lpAmount) public {
        require(balanceOf(msg.sender) >= lpAmount, "insufficient balance");

        address poolAddr = address(this);

        uint256 reserveTom = _tomToken.balanceOf(poolAddr);
        uint256 reserveJerry = _jerryToken.balanceOf(poolAddr);

        uint256 totalSupply = totalSupply();
        uint256 tomAmount = (lpAmount * reserveTom) / totalSupply;
        uint256 jerryAmount = (lpAmount * reserveJerry) / totalSupply;

        _burn(msg.sender, lpAmount);
        _tomToken.transfer(msg.sender, tomAmount);
        _jerryToken.transfer(msg.sender, jerryAmount);
    }

    function swapTomforJerry(uint256 tomIn, uint256 jerryOut) public {
        require(tomIn > 0, "insufficient balance");
        require(jerryOut > 0, "insufficient balance");

        address poolAddr = address(this);

        uint256 reserveTom = _tomToken.balanceOf(poolAddr);
        uint256 reserveJerry = _jerryToken.balanceOf(poolAddr);

        _tomToken.transferFrom(msg.sender, poolAddr, tomIn);
        _jerryToken.transfer(msg.sender, jerryOut);

        uint256 newReserveTom = _tomToken.balanceOf(poolAddr);
        uint256 newReserveJerry = _jerryToken.balanceOf(poolAddr);
        require(
            newReserveTom * newReserveJerry >= reserveTom * reserveJerry,
            "invalid k"
        );
    }

    function getTomAmountByJerry(uint256 jerryAmount)
        public
        view
        returns (uint256)
    {
        require(jerryAmount > 0, "amount should be more than 0");

        address poolAddr = address(this);

        uint256 reserveTom = _tomToken.balanceOf(poolAddr);
        uint256 reserveJerry = _jerryToken.balanceOf(poolAddr);
        require(reserveTom > 0 && reserveJerry > 0, "insufficient liquidity");

        return (jerryAmount * reserveTom) / reserveJerry;
    }

    function swapJerryforTom(uint256 jerryIn, uint256 tomOut) public {
        require(jerryIn > 0, "insufficient balance");
        require(tomOut > 0, "insufficient balance");

        address poolAddr = address(this);

        uint256 reserveTom = _tomToken.balanceOf(poolAddr);
        uint256 reserveJerry = _jerryToken.balanceOf(poolAddr);

        _tomToken.transfer(msg.sender, tomOut);
        _jerryToken.transferFrom(msg.sender, poolAddr, jerryIn);

        uint256 newReserveTom = _tomToken.balanceOf(poolAddr);
        uint256 newReserveJerry = _jerryToken.balanceOf(poolAddr);
        require(
            newReserveTom * newReserveJerry >= reserveTom * reserveJerry,
            "invalid k"
        );
    }

    function getJerryAmountByTom(uint256 tomAmount)
        public
        view
        returns (uint256)
    {
        require(tomAmount > 0, "amount should be more than 0");

        address poolAddr = address(this);

        uint256 reserveTom = _tomToken.balanceOf(poolAddr);
        uint256 reserveJerry = _jerryToken.balanceOf(poolAddr);
        require(reserveTom > 0 && reserveJerry > 0, "insufficient liquidity");

        return (tomAmount * reserveJerry) / reserveTom;
    }
}
