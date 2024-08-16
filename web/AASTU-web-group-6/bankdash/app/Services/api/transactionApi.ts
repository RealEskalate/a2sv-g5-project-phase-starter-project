import axios from 'axios';


const fetchTransaction = async () => {
    try {
        const response = await axios.get('https://bank-dashboard-6acc.onrender.com/transactions', {
            headers: {
                Authorization: `Bearer ${process.env.NAHOM_TOKEN}`,
            },
        });

        return response.data;
    } catch (error) {
        if (axios.isAxiosError(error)) {
            console.error('Axios error:', error.message);
        } else {
            // Handle non-Axios errors
            console.error('Unexpected error:', error);
        }
        throw error;
    }
};

export default fetchTransaction