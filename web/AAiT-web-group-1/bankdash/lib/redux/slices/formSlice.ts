import { createSlice } from "@reduxjs/toolkit";

const value = {name: "bjkbakjbvjb",
                email: "oiacnldkncal@pdojfaj.com",
                dateOfBirth: "2024-08-01",
                permanentAddress: "Addis abaaba",
                postalCode: "432424",
                username: "peofjpifjqpff",
                password: "pfjqpwfijepfj",
                presentAddress: "foiefnoif",
                city: "oeifhoqfih",
                country: "ofiehfoq",
                profilePicture: "foiofhofh",
                preference: {
                    currency: "dollar",
                    sentOrReceiveDigitalCurrency: true,
                    receiveMerchantOrder: true,
                    accountRecommendations: true,
                    timeZone: "east-african",
                    twoFactorAuthentication: true
}}



export const formSlice = createSlice({
    name: 'form',
    initialState:  { value: value },
    reducers: {
        setform: (state, newForm) => {
            state.value = {...state.value, ...newForm.payload}
            console.log(state.value)
        },
        setLastForm: (state, comingform) => {
            const newForm = {preference: comingform.payload}
            state.value = {...state.value, ...newForm}
            console.log(state.value)
        }
    }
})

export const  { setform, setLastForm } = formSlice.actions;
export default formSlice.reducer