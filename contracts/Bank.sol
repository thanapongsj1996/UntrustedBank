// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract Bank is Ownable, ERC20 {
    IERC20 private _token;
    mapping(address => uint256) _userBalance;

    constructor(ERC20 tokenAddr) ERC20("Bank", "Bank") {
        _token = IERC20(tokenAddr);
    }

    function deposit(uint256 amount) public {
        require(amount > 0, "deposit amount must more than zero");

        _userBalance[msg.sender] += amount;
        _token.transferFrom(msg.sender, address(this), amount);
    }

    function withdraw(uint256 amount) public {
        require(amount > 0, "withdraw amount must more than zero");
        require(
            _userBalance[msg.sender] >= amount,
            "you have not enough token to withdraw"
        );

        _userBalance[msg.sender] -= amount;
        _token.transfer(msg.sender, amount);
    }
}
