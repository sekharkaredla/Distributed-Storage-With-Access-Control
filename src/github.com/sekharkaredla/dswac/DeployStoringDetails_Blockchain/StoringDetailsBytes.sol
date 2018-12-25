pragma solidity >=0.4.22 <0.6.0;

contract StoringDetails{

    bytes public JWT;
    address owner;

    constructor(bytes memory initialJWT) public {
        owner = msg.sender;
        JWT = initialJWT;
    }

    function setJWT(bytes memory newJWT) public {
        JWT = newJWT;
    }
    function getJWT() public view returns (bytes memory){
        return JWT;
    }
}