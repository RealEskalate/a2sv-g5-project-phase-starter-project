// app/api/update-preference/route.js
import { NextResponse } from 'next/server';

const onSubmit = async (data, key) => {
  console.log(data, key, 'from onSubmit');

  try {
    const response = await fetch('https://bank-dashboard-6acc.onrender.com/user/update-preference', {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${key}`,
      },
      body: JSON.stringify(data),
    });

    if (!response.ok) {
      const errorText = await response.text();
      throw new Error(`Failed to update preferences: ${errorText}`);
    }

    const result = await response.json();
    console.log('Preferences updated successfully:', result);
    return NextResponse.json(result); // Return the result to the client
  } catch (error) {
    console.error('Error updating preferences:', error); // Log the error message
    return NextResponse.error(); // Return an error response
  }
};

export async function PUT(request) {
  try {
    const { data, key } = await request.json(); // Extract data and key from the request body
    return onSubmit(data, key); // Call onSubmit and return the response
  } catch (error) {
    console.error('Error processing PUT request:', error);
    return NextResponse.error(); // Return an error response
  }
}
