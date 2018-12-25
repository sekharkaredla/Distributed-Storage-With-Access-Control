pragma solidity >=0.4.22 <0.6.0;

contract StoringDetails{

    string public JWT;
    address owner;

    constructor(string memory initialJWT) public {
        owner = msg.sender;
        JWT = initialJWT;
    }

    function setJWT(string memory newJWT) public {
        JWT = newJWT;
    }
    function getJWT() public view returns (string memory){
        return JWT;
    }
}