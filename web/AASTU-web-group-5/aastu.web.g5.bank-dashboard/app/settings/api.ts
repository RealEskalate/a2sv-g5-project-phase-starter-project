  const onSubmit = async (data: any) => {
    const updatedData = {
      ...data,
      sentOrReceiveDigitalCurrency: digitalCurrency,
      receiveMerchantOrder: merchantOrder,
      accountRecommendations: accountRecommendations,
      twoFactorAuthentication: true,
    };

    try {
      const response = await axios.put('https://bank-dashboard-1tst.onrender.com/user/update-preference', updatedData, {
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${key}`,
        },
      });

      if (response.status === 200) {
        console.log('Preferences updated successfully:', response.data);
        dispatch(setUser(updatedData));
      } else {
        throw new Error(`Failed to update preferences: ${response.statusText}`);
      }
    } catch (error) {
      console.error('Error updating preferences:', error);
    }
  };