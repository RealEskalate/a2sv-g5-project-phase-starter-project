import React from 'react';
import InputGroup from './InputGroup';

const AddNewCardForm = () => {
  return (
    <form action=''>
      <div className='flex flex-col md:flex-row md:gap-6'>
        <InputGroup
          id='cardType'
          label='Card Type'
          inputType='text'
          registerName='cardType'
          register={undefined}
          placeholder='Classic'
        />
        <InputGroup
          id='cardHolder'
          label='Name On Card'
          inputType='text'
          registerName='cardHolder'
          register={undefined}
          placeholder='My Card'
        />
      </div>
      <div className='flex flex-col md:flex-row md:gap-6'>
        <InputGroup
          id='balance'
          label='Balance'
          inputType='text'
          registerName='balance'
          register={undefined}
          placeholder='27,00$'
        />
        <InputGroup
          id='expiryDate'
          label='Exipiration Date'
          inputType='text'
          registerName='expiryDate'
          register={undefined}
          placeholder='25 January 2025'
        />
      </div>
      <button
        type='submit'
        className='bg-[#1814f3] text-white px-10 py-3 rounded-lg w-full md:w-auto mt-4'
      >
        Add Card
      </button>
    </form>
  );
};

export default AddNewCardForm;
