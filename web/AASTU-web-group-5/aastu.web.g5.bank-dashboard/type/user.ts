
interface User {
    id: string;
    name: string;
    email: string;
    dateOfBirth: string;
    permanentAddress: string;
    postalCode: string;
    username: string;
    presentAddress: string;
    city: string;
    country: string;
    profilePicture: string;
    accountBalance: number;
    role: string;
    preference: {
        currency: string;
        sentOrReceiveDigitalCurrency: boolean;
        receiveMerchantOrder: boolean;
        accountRecommendations: boolean;
        timeZone: string;
        twoFactorAuthentication: boolean;
    };
}
 
export default User;