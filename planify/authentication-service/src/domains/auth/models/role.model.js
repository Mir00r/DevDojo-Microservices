module.exports = (sequelize, DataTypes) => {
    const Role = sequelize.define(
        "Role",
        {
            id: {
                type: DataTypes.INTEGER,
                autoIncrement: true,
                primaryKey: true,
            },
            name: {
                type: DataTypes.STRING,
                unique: true,
                allowNull: false,
            }
        },
        {schema: "auth"}
    );

    Role.associate = (models) => {
        Role.hasMany(models.User, {foreignKey: "roleId"});
    };

    return Role;
};
