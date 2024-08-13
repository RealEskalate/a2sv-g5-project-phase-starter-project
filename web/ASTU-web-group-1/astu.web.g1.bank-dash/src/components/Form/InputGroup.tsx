import { InputGroupType } from "@/types/formType";
const InputGroup = ({
  id,
  label,
  inputType,
  registerName,
  register,
  placeholder,
}: InputGroupType) => {
  return (
    <div className="w-6/12 space-y-3 my-3">
      <label htmlFor={id} className="gray-dark text-16px">
        {label} <br />
      </label>
      <input
        type={inputType}
        id={id}
        placeholder={placeholder}
        {...register(registerName)}
        className="w-full border-2 border-[#DFEAF2] p-5 py-3 rounded-xl placeholder:text-blue-steel focus:border-blue-steel outline-none"
      />
    </div>
  );
};

export default InputGroup;
