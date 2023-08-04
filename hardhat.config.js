require("@nomicfoundation/hardhat-toolbox");
require('dotenv').config();

/** @type import('hardhat/config').HardhatUserConfig */
module.exports = {
  solidity: "0.8.19",
  defaultNetwork:"goerli",
  networks:{
    hardhat:{},
    goerli:{
      url:process.env.goerliURL,
      accounts:[process.env.privateKey],
    }
  }
};
