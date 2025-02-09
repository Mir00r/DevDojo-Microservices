const nodemailer = require('nodemailer');

class EmailService {
    constructor() {
        this.transporter = nodemailer.createTransport({
            host: process.env.SMTP_HOST,
            port: process.env.SMTP_PORT,
            secure: process.env.SMTP_SECURE === 'true',
            auth: {
                user: process.env.SMTP_USER,
                pass: process.env.SMTP_PASS
            }
        });
    }

    async sendVerificationEmail(user, token) {
        const verificationUrl = `${process.env.FRONTEND_URL}/verify-email?token=${token}`;

        const mailOptions = {
            from: process.env.SMTP_FROM,
            to: user.email,
            subject: 'Email Verification',
            html: `
                <h1>Verify Your Email</h1>
                <p>Please click the link below to verify your email address:</p>
                <a href="${verificationUrl}">${verificationUrl}</a>
                <p>This link will expire in 24 hours.</p>
            `
        };

        try {
            await this.transporter.sendMail(mailOptions);
        } catch (error) {
            console.error('Email sending failed:', error);
            // Don't throw error - we don't want registration to fail if email fails
        }
    }

    async sendPasswordResetEmail(email, token) {
        const resetUrl = `${process.env.FRONTEND_URL}/reset-password?token=${token}`;

        const mailOptions = {
            from: process.env.SMTP_FROM,
            to: email,
            subject: 'Password Reset',
            html: `
                <h1>Reset Your Password</h1>
                <p>You requested a password reset. Click the link below to reset your password:</p>
                <a href="${resetUrl}">${resetUrl}</a>
                <p>This link will expire in 1 hour.</p>
                <p>If you didn't request this, please ignore this email.</p>
            `
        };

        try {
            await this.transporter.sendMail(mailOptions);
        } catch (error) {
            console.error('Password reset email sending failed:', error);
            throw new Error('Failed to send password reset email');
        }
    }
}

module.exports = new EmailService();
