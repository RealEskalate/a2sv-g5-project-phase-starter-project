import '../../domain/entities/log_in.dart';
import '../../domain/entities/sign_up.dart';
import 'log_in_model.dart';
import 'sign_up_model.dart';

extension LogInMapper on LogInEntity {
  LogInModel toProductModel() {
    return LogInModel(
      email: email,
      password: password,
    );
  }
}

extension SignUpMapper on SignUpEntity {
  SignUpModel toSignUpModel() {
    return SignUpModel(
      email: email,
      password: password,
      username: username,
    );
  }
}
