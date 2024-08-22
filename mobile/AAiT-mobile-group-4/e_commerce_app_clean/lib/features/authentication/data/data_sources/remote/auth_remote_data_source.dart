import '../../model/log_in_model.dart';
import '../../model/sign_up_model.dart';
import '../../model/user_model.dart';

abstract class AuthRemoteDataSource {
  Future<void> signUp(SignUpModel signUpModel);
  Future<void> logIn(LogInModel logInModel);
  Future<void> logOut();
  Future<UserModel> getCurrentUser();
}