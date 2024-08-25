part of 'user_input_validation_cubit.dart';

sealed class UserInputValidationState extends Equatable {
  final String name,
      email,
      password,
      confirmPassword,
      checkbox,
      passwordContent;
  const UserInputValidationState(
      {this.name = AppData.strInitial,
      this.email = AppData.strInitial,
      this.password = AppData.strInitial,
      this.confirmPassword = AppData.strInitial,
      this.checkbox = AppData.strInitial,
      this.passwordContent = ''});

  String? validate();

  @override
  List<Object> get props => [name, email, password, confirmPassword, checkbox];
}

final class UserInputValidationInitial extends UserInputValidationState {
  @override
  String? validate() {
    return 'Fill the fields';
  }
}

final class LoginUserInputValidated extends UserInputValidationState {
  @override
  // ignore: overridden_fields
  final String email;
  @override
  // ignore: overridden_fields
  final String password;
  @override
  // ignore: overridden_fields
  final String checkbox;
  final List<String> errorMessages;

  const LoginUserInputValidated(
      {required this.email,
      required this.password,
      required this.checkbox,
      this.errorMessages = const <String>[]})
      : super();
  @override
  String? validate() {
    if (email == AppData.strNotValidated || email == AppData.strInitial) {
      return AppData.getMessage(AppData.invalidEmail);
    }
    if (password == AppData.strNotValidated || password == AppData.strInitial) {
      return AppData.getMessage(AppData.invalidPassword);
    }
    return null;
  }

  bool getSingleInputState(String type) {
    Map<String, String> map = {
      InputDataValidator.email: email,
      InputDataValidator.password: password,
      InputDataValidator.checkBox: checkbox
    };
    return map[type] == AppData.strValidated || map[type] == AppData.strInitial;
  }
}

final class SignupUserInputValidated extends UserInputValidationState {
  @override
  // ignore: overridden_fields
  final String email;
  @override
  // ignore: overridden_fields
  final String password;
  @override
  // ignore: overridden_fields
  final String name;
  @override
  // ignore: overridden_fields
  final String checkbox;

  @override
  // ignore: overridden_fields
  final String confirmPassword;
  @override
  // ignore: overridden_fields
  final String passwordContent;

  final List<String> errorMessages;

  const SignupUserInputValidated(
      {required this.name,
      required this.email,
      required this.password,
      required this.checkbox,
      required this.confirmPassword,
      required this.passwordContent,
      this.errorMessages = const <String>[]})
      : super(
            name: name,
            email: email,
            password: password,
            checkbox: checkbox,
            passwordContent: passwordContent);

  @override
  String? validate() {
    if (name == AppData.strNotValidated || name == AppData.strInitial) {
      return AppData.getMessage(AppData.invalidName);
    }
    if (email == AppData.strNotValidated || email == AppData.strInitial) {
      return AppData.getMessage(AppData.invalidEmail);
    }
    if (password == AppData.strNotValidated || password == AppData.strInitial) {
      return AppData.getMessage(AppData.invalidPassword);
    }
    // if (confirmPassword == AppData.strNotValidated ||
    //     confirmPassword == AppData.strInitial) {
    //   return AppData.getMessage(AppData.confirmPassword);
    // }
    if (checkbox == AppData.strNotValidated || checkbox == AppData.strInitial) {
      return AppData.getMessage(AppData.checkbox);
    }

    return null;
  }

  bool getSingleInputState(String type) {
    Map<String, String> map = {
      InputDataValidator.name: name,
      InputDataValidator.email: email,
      InputDataValidator.password: password,
      InputDataValidator.confirmPass: confirmPassword,
      InputDataValidator.checkBox: checkbox
    };
    return map[type] == AppData.strValidated || map[type] == AppData.strInitial;
  }
}
