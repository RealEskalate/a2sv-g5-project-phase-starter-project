import { useState } from "react";

interface Step3Props {
  data: { [key: string]: any };
  onNext: (data: { [key: string]: any }) => void;
  onPrevious: () => void;
}

const Step3: React.FC<Step3Props> = ({ data, onNext, onPrevious }) => {
  const [formData, setFormData] = useState<{ [key: string]: any }>(data);
  console.log("Step 3 Data:", formData);

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
  };

  return (
    <div>
      <h2>Step 3</h2>
      <input
        type="text"
        name="field3"
        value={formData.field3 || ""}
        onChange={handleChange}
        className="border p-2"
        placeholder="Field 3"
      />
      <div className="flex justify-between mt-4">
        <button onClick={onPrevious} className="px-4 py-2 bg-gray-200 rounded">
          Previous
        </button>
        <button
          onClick={() => onNext(formData)}
          className="px-4 py-2 bg-blue-500 text-white rounded"
        >
          Save and Submit
        </button>
      </div>
    </div>
  );
};

export default Step3;
