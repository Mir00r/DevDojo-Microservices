const {catchAsync} = require('../../../utils/errorHandler');
const MfaService = require('../services/mfaService');
const {ApiResponse} = require('../../../utils/apiResponse');

class MfaController {
    /**
     * Enable MFA for user
     */
    enableMfa = catchAsync(async (req, res) => {
        const result = await MfaService.enableMfa({
            userId: req.user.id,
            email: req.user.email
        });

        return ApiResponse.success(res, {
            message: 'MFA setup initiated successfully',
            data: result
        });
    });

    /**
     * Verify MFA code
     */
    verifyMfa = catchAsync(async (req, res) => {
        const result = await MfaService.verifyMfa({
            userId: req.user.id,
            code: req.body.code
        });

        return ApiResponse.success(res, {
            message: result.message,
            data: {verified: result.verified}
        });
    });
}

module.exports = new MfaController();
