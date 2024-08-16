import axios, { isAxiosError } from 'axios';
import UserValue from '@/types/UserValue';
import ResponseValue from '@/types/ResponseValue';
import SignupResponseValue from '@/types/SignupResponseValue';


const SignupService = async (formData: UserValue): Promise<ResponseValue> => {
    try {
        const response = await axios.post<SignupResponseValue>(
            "https://bank-dashboard-6acc.onrender.com/auth/register",
            formData,
            {
                headers: {
                    "Content-Type": "application/json",
                },
            }
        );

        if (response.status === 200) {
            return { success: true, data: response.data };
        } else {
            return { success: false, data: null };
        }
    } catch (err) {
        if (isAxiosError(err) && err.response) {
            return { success: false, data: null };
        } else {
            return { success: false, data: null };
        }
    }
};

export default SignupService;
