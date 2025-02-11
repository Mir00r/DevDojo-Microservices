// HTTP Status codes as constants
const HttpStatus = {
    OK: 200,
    CREATED: 201,
    NO_CONTENT: 204,
    BAD_REQUEST: 400,
    UNAUTHORIZED: 401,
    FORBIDDEN: 403,
    NOT_FOUND: 404,
    CONFLICT: 409,
    INTERNAL_SERVER: 500
};

// Response messages
const ResponseMessage = {
    SUCCESS: 'success',
    ERROR: 'error',
    FAIL: 'fail'
};

class ApiResponse {
    /**
     * Success Response
     * @param {Object} res - Express response object
     * @param {number} statusCode - HTTP status code
     * @param {string} message - Success message
     * @param {Object} data - Response data
     */
    static success(res, {statusCode = HttpStatus.OK, message = '', data = null}) {
        return res.status(statusCode).json({
            status: ResponseMessage.SUCCESS,
            ...(message && {message}),
            ...(data && {data})
        });
    }

    /**
     * Error Response
     * @param {Object} res - Express response object
     * @param {number} statusCode - HTTP status code
     * @param {string} message - Error message
     * @param {Object} errors - Error details
     */
    static error(res, {statusCode = HttpStatus.INTERNAL_SERVER, message = 'Internal Server Error', errors = null}) {
        return res.status(statusCode).json({
            status: ResponseMessage.ERROR,
            message,
            ...(errors && {errors})
        });
    }

    /**
     * Created Response
     * @param {Object} res - Express response object
     * @param {string} message - Success message
     * @param {Object} data - Created resource data
     */
    static created(res, {message = 'Resource created successfully', data = null}) {
        return this.success(res, {
            statusCode: HttpStatus.CREATED,
            message,
            data
        });
    }

    /**
     * No Content Response
     * @param {Object} res - Express response object
     */
    static noContent(res) {
        return res.status(HttpStatus.NO_CONTENT).send();
    }
}

module.exports = {
    HttpStatus,
    ResponseMessage,
    ApiResponse
};
