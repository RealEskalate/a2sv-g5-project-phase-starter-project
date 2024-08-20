import { getAllCards } from '@/services/cardfetch'

export const useFetchCards = (token: string) => {    
    const fetchCards = async () => {
        try {
            const result = await getAllCards(token);
            const data = result.content
            return data;
        } catch (err) {
            console.error("Failed to get data",err)
        }
    };
    
    return fetchCards();
}
