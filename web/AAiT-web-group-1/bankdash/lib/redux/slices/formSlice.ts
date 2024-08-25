import { createSlice } from "@reduxjs/toolkit";

const value = {name: "bjkbafdsafkjbvjb",
                email: "nferefdddsdcafefefl@pdojfaj.com",
                dateOfBirth: "2024-08-01",
                permanentAddress: "Addis abaaba",
                postalCode: "432424",
                username: "dpiffddffefefcjf",
                password: "pfjfefedsdffdfefqpwfcijepfj",
                presentAddress: "foieffdnoif",
                city: "oeifhoqfih",
                country: "ofiehfoq",
                profilePicture: "foiofhfdofh",
                preference: {
                    currency: "birr",
                    sentOrReceiveDigitalCurrency: true,
                    receiveMerchantOrder: true,
                    accountRecommendations: true,
                    timeZone: "opi",
                    twoFactorAuthentication: true
}}



export const formSlice = createSlice({
    name: 'form',
    initialState:  { value: value },
    reducers: {
        setform: (state, newForm) => {
            state.value = {...state.value, ...newForm.payload}
        },
    }
})

export const  { setform } = formSlice.actions;
export default formSlice.reducer