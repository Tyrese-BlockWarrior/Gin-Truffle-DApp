const Cert = artifacts.require("Cert");
const { writeFileSync } = require("fs");

module.exports = function (_deployer) {
  // Use deployer to state migration tasks.
  _deployer
    .deploy(Cert)
    .then((instance) => {
      let details = instance.constructor.class_defaults;
      details["contract"] = instance.address;
      writeFileSync("details.json", JSON.stringify(details, null, 2));
      writeFileSync("Cert.abi", JSON.stringify(instance.abi, null, 2));
    })
    .catch((err) => {
      console.log(err);
    });
};
