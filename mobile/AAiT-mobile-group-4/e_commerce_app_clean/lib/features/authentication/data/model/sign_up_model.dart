import '../../domain/entities/sign_up.dart';

class SignUpModel extends SignUpEntity {
  const SignUpModel({
    required super.email,
    required super.password,
    required super.username,
  });

  Map<String,dynamic> toJson() {
    return {
      'email': email,
      'password': password,
      'name': username,
    };
  }
}
