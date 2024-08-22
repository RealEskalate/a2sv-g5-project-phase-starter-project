import '../../domain/entities/register_entity.dart';

class RegisterModel extends RegistrationEntity {
  RegisterModel(
      {required super.name, required super.email, required super.password});

  Map<String, dynamic> toJson() {
    return {'name': name, 'email': email, 'password': password};
  }

  static RegisterModel toModel(RegistrationEntity registrationEntity) {
    return RegisterModel(
        name: registrationEntity.name,
        email: registrationEntity.email,
        password: registrationEntity.password);
  }
}
