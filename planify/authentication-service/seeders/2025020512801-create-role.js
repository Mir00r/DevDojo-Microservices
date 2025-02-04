module.exports = {
    up: async (queryInterface) => {
        return queryInterface.bulkInsert({schema: "auth", tableName: "roles"}, [{name: "Admin"}, {name: "User"}]);
    }, down: async (queryInterface) => {
        return queryInterface.bulkDelete({schema: "auth", tableName: "roles"}, null, {});
    },
};
