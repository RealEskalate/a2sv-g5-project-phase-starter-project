import '../../domain/entities/sign_up_entity.dart';

class SignUpModel extends SignUpEntity {
  const SignUpModel(
      {required super.name,
      required super.email,
      required super.password,
      required super.repeatedPassword});
}
