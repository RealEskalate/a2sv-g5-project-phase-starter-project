import { User } from '@/types/User';
import { Preference } from '@/types/User';

export interface UserSignup {
    name: string;
    email: string;
    dateOfBirth: string; 
    permanentAddress: string;
    postalCode: string;
    username: string;
    password: string;
    presentAddress: string;
    city: string;
    country: string;
    preference: Preference;
}