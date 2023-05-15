const Cert = artifacts.require("Cert");
const { writeFileSync } = require("fs");

module.exports = function (_deployer) {
  // Use deployer to state migration tasks.
  _deployer
    .deploy(Cert)
    .then((instance) => {
      writeFileSync(
        "deployer.json",
        JSON.stringify(instance.constructor.class_defaults, null, 2)
      );
      writeFileSync(
        "Cert.abi",
        JSON.stringify(instance.abi, null, 2)
      );
    })
    .catch((err) => {
      console.log(err);
    });
};
