import '../entity/user.dart';
import '../repository/authentication_repository.dart';

class SignUpUseCase{
  final AuthenticationRepository repository;
  SignUpUseCase(this.repository);

  Future<User> call({required String email, required String password, required String username}) async {
    try{
      print("from usecase");
    return await repository.signUp(email: email, password: password, username: username);}
    catch(e){
      throw Exception('Server Failure');
    }
  }
}