import '../entity/user.dart';
import '../repository/authentication_repository.dart';

class SignUpUseCase{
  final AuthenticationRepository repository;
  SignUpUseCase(this.repository);

  Future<User> call({required String email, required String password, required String username}) async {
    try{
      print("from usecase: attempting to sign up user");
    final user = await repository.signUp(email: email, password: password, username: username);
     print("from usecase: user signed up successfully - $user");
     return user;
    }
    catch(e){
      print("from usecase: error during sign up - $e");
      throw Exception('Server Failure: ${e.toString()}');

    }
  }
}