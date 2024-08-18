import InputGroup from "../InputGroup";

const SignUpForm = () => {
  return (
    <form className="flex flex-col items-center w-full md:w-1/2 justify-center md:px-16">
      <div className="w-full md:flex md:gap-4 ">
        <InputGroup
          id="name"
          label="Full Name"
          inputType="name"
          registerName="name"
          register=""
          placeholder="Enter Full Name"
        />
        <InputGroup
          id="email"
          label="Email"
          inputType="email"
          registerName="email"
          register=""
          placeholder="Enter Your Email"
        />
      </div>

      <div className="w-full md:flex md:gap-4 ">
        <InputGroup
          id="dateOfBirth"
          label="Date Of Birth"
          inputType="date"
          registerName="dateOfBirth"
          register=""
          placeholder="Enter Date Of Birth"
        />
        <InputGroup
          id="permanentAddress"
          label="Permanent Address"
          inputType="text"
          registerName="permanentAddress"
          register=""
          placeholder="Enter Permanent Address"
        />
      </div>

      <div className="w-full md:flex md:gap-4 ">
        <InputGroup
          id="postalCode"
          label="Postal Code"
          inputType="text"
          registerName="postalCode"
          register=""
          placeholder="Enter Postal Code"
        />
        <InputGroup
          id="username"
          label="Username"
          inputType="username"
          registerName="username"
          register=""
          placeholder="Enter Username"
        />
      </div>
      <div className="w-full md:flex md:gap-4 ">
        <InputGroup
          id="password"
          label="password"
          inputType="password"
          registerName="password"
          register=""
          placeholder="Enter Password"
        />
        <InputGroup
          id="password"
          label="Confirm Password"
          inputType="password"
          registerName="password"
          register=""
          placeholder="RE-Enter password"
        />
      </div>
      <div className="w-full md:flex md:gap-4 ">
        <InputGroup
          id="presentAddress"
          label="Present Address"
          inputType="text"
          registerName="presentAddress"
          register=""
          placeholder="Enter Present Address"
        />
        <InputGroup
          id="city"
          label="City"
          inputType="text"
          registerName="city"
          register=""
          placeholder="Enter City"
        />
      </div>

      <div className="w-full md:flex md:gap-4 ">
        <InputGroup
          id="country"
          label="Country"
          inputType="text"
          registerName="country"
          register=""
          placeholder="Enter Country"
        />
        <InputGroup
          id="profilePicture"
          label="profilePicture"
          inputType="file"
          registerName="profilePicture"
          register=""
          placeholder="Enter Profile Picture"
        />
      </div>

      <div className="w-full md:flex md:gap-4 ">
        <div className=" w-full lg:w-6/12 space-y-3 my-3">
          <label htmlFor="timeZone" className="gray-dark text-16px">
            Currency
          </label>
          <select
            id="timeZone"
            className="w-full border-2 border-[#DFEAF2] p-5 py-3 rounded-xl placeholder:text-blue-steel focus:border-blue-steel outline-none"
          >
            <option selected>USD</option>
            <option value="US">Birr</option>
            <option value="CA">Birr</option>
            <option value="FR">Birr</option>
          </select>
        </div>
        <div className=" w-full lg:w-6/12 space-y-3 my-3">
          <label htmlFor="timeZone" className="gray-dark text-16px">
            Time Zone
          </label>
          <select
            id="timeZone"
            className="w-full border-2 border-[#DFEAF2] p-5 py-3 rounded-xl placeholder:text-blue-steel focus:border-blue-steel outline-none"
          >
            <option selected>GMT 3+</option>
            <option value="US">Birr</option>
            <option value="CA">Birr</option>
            <option value="FR">Birr</option>
          </select>
        </div>
      </div>

      <button
        type="submit"
        className="bg-[#1814f3] text-white px-10 py-3 font-Lato font-bold rounded-lg w-full mt-4"
      >
        Login
      </button>
    </form>
  );
};

export default SignUpForm;
