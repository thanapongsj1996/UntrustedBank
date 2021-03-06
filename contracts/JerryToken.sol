// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract JerryToken is ERC20 {
    constructor() ERC20("Token", "Token") {}

    function mint(uint256 amount) public {
        _mint(msg.sender, amount);
    }
}
