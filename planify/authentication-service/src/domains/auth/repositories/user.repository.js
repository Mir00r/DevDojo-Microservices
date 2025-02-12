// repositories/user.repository.js
const BaseRepository = require('./base.repository');
const {User, Role} = require('../../../../models');
const {Op} = require("sequelize");

class UserRepository extends BaseRepository {
    constructor() {
        super(User);
    }

    async findByEmail(email) {
        return this.findOne({
            where: {email: email},
            include: [{
                model: Role,
                as: 'role',
                attributes: ['name']
            }]
        });
    }

    async findAllWithPagination({page = 1, limit = 10, search = '', roleId = null}) {
        const offset = (page - 1) * limit;
        const whereClause = {};

        if (search) {
            whereClause[Op.or] = [
                {name: {[Op.iLike]: `%${search}%`}},
                {email: {[Op.iLike]: `%${search}%`}}
            ];
        }

        if (roleId) {
            whereClause.roleId = roleId;
        }

        return this.model.findAndCountAll({
            where: whereClause,
            include: [{
                model: Role,
                as: 'role',
                attributes: ['name']
            }],
            limit,
            offset,
            order: [['createdAt', 'DESC']]
        });
    }

    async updateLastLogin(userId) {
        return this.update(userId, {lastLogin: new Date()});
    }
}

module.exports = new UserRepository();
