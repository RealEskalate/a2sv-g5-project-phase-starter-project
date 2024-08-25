class AppData {
  //! Shared Preference list name
  static const String sharedProduct = 'my_shared_products';
  static const String tokenPlacement = 'token_placement';
  //! Url data's
  static const String baseUrl =
      'https://g5-flutter-learning-path-be.onrender.com/api/v1';
  static const String baseUrlV2 =
      'https://g5-flutter-learning-path-be.onrender.com/api/v2';
  static const String registerUser = '$baseUrlV2/auth/register';
  static const String logInUser = '$baseUrlV2/auth/login';
  static const String allProductUrl = '$baseUrlV2/products';
  //! json information
  static const Map<String, String> jsonHeader = {
    'Content-Type': 'application/json'
  };

  //! request methids
  static const String post = 'POST';
  static const String get = 'GET';
  static const String put = 'PUT';
  static const String delete = 'DELETE';

  //! validation data
  static const String strValidated = 'VALIDATED';
  static const String strInitial = 'INITIAL';
  static const String strNotValidated = 'NOT_VALIDATED';

  /// The following are error messages

  static const int cacheError = 1001;
  static const int serverError = 1002;
  static const int connectionError = 1003;
  static const int successInsert = 1004;
  static const int successUpdate = 1005;
  static const int successDelete = 1006;
  static const int userExists = 1007;
  static const int loginFailed = 1008;
  //! validation errors
  static const int invalidPriceCharacter = 2000;
  static const int negativePrice = 2001;
  static const int invalidName = 2002;
  static const int invalidPassword = 2003;
  static const int invalidEmail = 2004;
  static const int checkbox = 2005;
  static const int confirmPassword = 2006;
  static const int logoutError = 2007;

  static const Map<int, String> message = {
    cacheError: 'Caching Failed',
    connectionError: 'No Internet Connection',
    serverError: 'Server request failed',
    userExists: 'User Already Exists',
    loginFailed: 'Login failed, check your email and password!',
    successDelete: 'Successfully Deleted',
    successInsert: 'Successfully inserted',
    successUpdate: 'Successfully Update',
    invalidPriceCharacter: 'Price can not contain character',
    negativePrice: 'Price can not be negative',
    invalidName: 'Check the name field!',
    invalidEmail: 'Check the email field!',
    invalidPassword: 'Password length must be 8 or greater',
    checkbox: 'Must agree to our term and policy',
    confirmPassword: 'Your passwords must be similar',
    logoutError: 'Couldn\'t logout please try again!'
  };

  static const Map<String, int> methodInt = {
    post: successInsert,
    put: successUpdate,
    delete: successDelete
  };

  static int getCorrespondingSuccess(String method) {
    return methodInt[method]!;
  }

  static String getMessage(int code) {
    return message[code]!;
  }

  static const String imageUrl =
      'https://static.nike.com/a/images/f_auto/dpr_3.0,cs_srgb/w_363,c_limit/83e721fb-4f3f-44d5-ae11-5ef19006fd93/best-running-shoes-for-walking-by-nike.jpg';

  //!these are page constants
  static const String login = 'L';
  static const String signup = 'S';
}
