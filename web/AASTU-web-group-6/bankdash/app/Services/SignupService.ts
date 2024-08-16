import axios, { isAxiosError } from 'axios';
import UserValue from '@/types/UserValue';
import ResponseValue from '@/types/ResponseValue';


const SignupService = async (formData: UserValue): Promise<ResponseValue> => {
    try {
        const response = await axios.post(
            "https://bank-dashboard-6acc.onrender.com/auth/signup",
            formData,
            {
                headers: {
                    "Content-Type": "application/json",
                },
            }
        );

        if (response.status === 200) {
            return { success: true, data: response };
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
