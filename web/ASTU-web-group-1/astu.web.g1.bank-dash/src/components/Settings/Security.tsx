import InputGroup from "../Form/InputGroup";
import ToggleInput from "../Form/ToggleInput";

const Titles = ({ title }: { title: string }) => {
  return <h2 className="text-17px font-semibold text-[#333b69]">{title}</h2>;
};

const Security = () => {
  return (
    <div>
      <Titles title="Two-factor Authentication" />

      <ToggleInput
        label="I send or receive digital currency"
        inputType="checkbox"
        id="email"
        registerName="email"
        register={undefined}
        placeholder="Email"
        currentState={true}
      />

      <Titles title="Change Password" />

      <InputGroup
        id="password"
        label="Current Password"
        inputType="password"
        registerName="password"
        register={undefined}
        placeholder="*********************"
      />
      <InputGroup
        id="newPassword"
        label="New Password"
        inputType="password"
        registerName="newPassword"
        register={undefined}
        placeholder="*********************"
      />

      <div className="flex justify-end">
        <button
          type="submit"
          className="bg-[#1814f3] text-white px-10 py-2 rounded-lg w-full md:w-auto mt-4"
        >
          Submit
        </button>
      </div>
    </div>
  );
};

export default Security;
