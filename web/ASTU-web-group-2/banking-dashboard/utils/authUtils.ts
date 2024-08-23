import {jwtDecode} from 'jwt-decode';

// Function to check if a JWT is expired
export function isTokenExpired(token:string) {
    try {
        const decoded = jwtDecode(token);
        if (!decoded || !decoded.exp) {
            return true; 
        }

        const currentTime = Date.now() / 1000; // Current time in second
        return decoded.exp < currentTime; // Check if the token has expired
    } catch (error) {
        return true; // Consider the token expired in case of an error
    }
}
