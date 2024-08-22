import { useState } from "react";

interface Step2Props {
  data: { [key: string]: any };
  onNext: (data: { [key: string]: any }) => void;
  onPrevious: () => void;
}

const Step2: React.FC<Step2Props> = ({ data, onNext, onPrevious }) => {
  const [formData, setFormData] = useState<{ [key: string]: any }>(data);

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
  };

  return (
    <div>
      <h2>Step 2</h2>
      <input
        type="text"
        name="field2"
        value={formData.field2 || ""}
        onChange={handleChange}
        className="border p-2"
        placeholder="Field 2"
      />
      <div className="flex justify-between mt-4">
        <button onClick={onPrevious} className="px-4 py-2 bg-gray-200 rounded">
          Previous
        </button>
        <button
          onClick={() => onNext(formData)}
          className="px-4 py-2 bg-blue-500 text-white rounded"
        >
          Save and Next
        </button>
      </div>
    </div>
  );
};

export default Step2;
