import { InputGroupType } from "@/types/formType";
const InputGroup = ({
  id,
  label,
  inputType,
  registerName,
  register,
  placeholder,
  errorMessage,
  min,
}: InputGroupType) => {
  return (
    <div className="w-full space-y-1 my-3">
      <label htmlFor={id} className="gray-dark text-16px">
        {label} <br />
      </label>
      <input
        type={inputType}
        min={min}
        id={id}
        placeholder={placeholder}
        {...register(registerName)}
        className="w-full border-2 border-[#DFEAF2] p-5 py-3 rounded-xl placeholder:text-blue-steel focus:border-blue-steel outline-none"
      />
      {errorMessage && <p className="text-red-400"> {errorMessage} </p>}
    </div>
  );
};

export default InputGroup;
