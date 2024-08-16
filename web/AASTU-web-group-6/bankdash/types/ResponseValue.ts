import SignupResponseValue from "./SignupResponseValue"
interface ResponseValue {
    success : boolean
    data: SignupResponseValue | null
}

export default ResponseValue