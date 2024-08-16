import InputGroup from "./InputGroup"

const EditProfileForm = () => {
  return (
    <div className="w-full">
        <form action="">
            <div className="flex flex-col md:flex-row md:space-x-5">
                <InputGroup id="name" label="Your Name" inputType="text" registerName="name" register={undefined} placeholder="Charlene Reed" />
                <InputGroup id="username" label="User Name" inputType="text" registerName="username" register={undefined} placeholder="Charlene Reed" />
            </div>
            <div className="flex flex-col md:flex-row md:space-x-5">
                <InputGroup id="email" label="Email" inputType="text" registerName="email" register={undefined} placeholder="charlene.reed@gmail.com" />
                <InputGroup id="password" label="Password" inputType="password" registerName="password" register={undefined} placeholder="*********************" />
            </div>
            <div className="flex flex-col md:flex-row md:space-x-5">
                <InputGroup id="dateOfBirth" label="Date Of Birth" inputType="date" registerName="dateOfBirth" register={undefined} placeholder="25 January 1990" />
                <InputGroup id="presentAddress" label="Present Address" inputType="text" registerName="presentAddress" register={undefined} placeholder="San Jose, California, USA" />
            </div>
            <div className="flex flex-col md:flex-row md:space-x-5">
                <InputGroup id="permanentAddress" label="Permanent Address" inputType="text" registerName="permanentAddress" register={undefined} placeholder="San Jose, California, USA" />
                <InputGroup id="city" label="City" inputType="text" registerName="city" register={undefined} placeholder="San Jose" />
            </div>
            <div className="flex flex-col md:flex-row md:space-x-5">
                <InputGroup id="postalCode" label="Postal Code" inputType="text" registerName="postalCode" register={undefined} placeholder="45322" />
                <InputGroup id="country" label="Country" inputType="text" registerName="country" register={undefined} placeholder="USA" />
            </div>
            <div className="flex justify-end"><button type="submit" className="bg-[#1814f3] text-white px-10 py-2 rounded-lg w-full md:w-auto mt-4">Submit</button></div>
        </form>
    </div>
  )
}

export default EditProfileForm
