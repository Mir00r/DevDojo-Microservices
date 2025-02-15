const {PrivilegeResponseDto} = require('../dtos/privilege.dto');
const PrivilegeRepository = require('../repositories/privilegeRepository');

class PrivilegeService {
    async createPrivilege(createDto) {
        const privilege = await PrivilegeRepository.create(createDto);
        return PrivilegeResponseDto.from(privilege);
    }

    async getAllPrivileges(queryDto) {
        const privileges = await PrivilegeRepository.findAll(queryDto);
        return PrivilegeResponseDto.fromPaginated(privileges);
    }

    async getPrivilegeById(id) {
        const privilege = await PrivilegeRepository.findById(id);
        return PrivilegeResponseDto.from(privilege);
    }

    async updatePrivilege(id, updateDto) {
        const privilege = await PrivilegeRepository.update(id, updateDto);
        return PrivilegeResponseDto.from(privilege);
    }

    async deletePrivilege(id) {
        await PrivilegeRepository.delete(id);
    }
}

module.exports = new PrivilegeService();