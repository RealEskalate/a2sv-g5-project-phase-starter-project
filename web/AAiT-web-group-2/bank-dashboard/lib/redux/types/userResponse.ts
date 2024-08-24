import { User } from "@/types/User";

export interface UserResponse{
    success: boolean;
    message: string;
    data: {
        access_token: string;
        refresh_token: string;
        data: User;
    }
}