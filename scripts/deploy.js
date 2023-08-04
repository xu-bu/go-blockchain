const { ethers } = require("hardhat");

async function main() {
  const contractFactory = await ethers.getContractFactory("NFTMarket")
  const myContract = await contractFactory.deploy()
  console.log("deploy successfully, address:", myContract.target)
}

main()
  .then(() => process.exit(0))
  .catch((e) => {
    console.error(e);
    process.exit(1);
  });
