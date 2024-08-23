import ProfileUpload from "@/app/dashboard/_components/ProfileUpload";
import { SignUpForm } from "./signUpForm";
// import Stepper from "./Stepper";

const SignUp = async () => {
  return (
    <section className="flex items-center justify-center size-full max-sm:px-6">
      <SignUpForm />
      {/* <Stepper /> */}
    </section>
  );
};

export default SignUp;
