import { ToggleInputType } from "@/types/formType";

const ToggleInput = ({
  label,
  inputType,
  id,
  register,
  placeholder,
  registerName,
  currentState
}: ToggleInputType) => {
  return (
    <div className="bloack my-4">
      <label className="inline-flex items-center cursor-pointer">
        {/* <input type="checkbox" value="" className="sr-only peer" /> */}
        <input
          type={inputType}
          id={id}
          placeholder={placeholder}
          // {...register(registerName)}
          className="sr-only peer"
        />
        <div className={`${currentState ? 'bg-[#16dbcc]' : 'bg-gray-200'} relative w-11 h-6 peer-focus:outline-none peer-focus:ring-2 peer-focus:ring-[#16dbcc] rounded-full peer 
        peer-checked:after:translate-x-full rtl:peer-checked:after:-translate-x-full peer-checked:after:border-white after:content-[''] 
        after:absolute after:top-[2px] after:start-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full 
        after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600`}></div>

        <span className="ms-3 text-sm font-medium text-gray-900">
          {label}
        </span>
      </label>
    </div>
  );
};

export default ToggleInput;
