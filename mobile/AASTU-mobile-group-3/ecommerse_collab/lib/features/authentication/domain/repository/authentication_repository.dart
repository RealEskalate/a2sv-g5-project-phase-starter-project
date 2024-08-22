import '../entity/user.dart';

abstract class AuthenticationRepository{
  Future<User> signUp({String email, String password, String username});
  Future<void> logOut();
  Future<User> logIn({String email, String password});

  Future<bool> isSignedIn();

}