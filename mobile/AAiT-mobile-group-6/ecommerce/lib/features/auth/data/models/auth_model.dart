import '../../domain/entities/log_in_entity.dart';
import '../../domain/entities/sign_up_entity.dart';

class SignUPModel extends SignUpEntity {
  const SignUPModel({
    required super.id,
    required super.name,
    required super.email,
    required super.password,
  });

  factory SignUPModel.fromJson(Map<String, dynamic> json) {
    return SignUPModel(
        id: json['_id'],
        name: json['name'],
        email: json['email'],
        password: json['password']);
  }

  Map<String, dynamic> toJson() {
    return {
      'id': super.id,
      'name': super.name,
      'email': super.email,
      'password': super.password,
    };
  }
}

class LogInModel extends LogInEntity {
  const LogInModel({
    required super.id,
    required super.email,
    required super.password,
  });

  factory LogInModel.fromJson(Map<String, dynamic> json) {
    return LogInModel(
        id: json['_id'], email: json['email'], password: json['password']);
  }

  Map<String, dynamic> toJson() {
    return {
      'id': super.id,
      'email': super.email,
      'password': super.password,
    };
  }
}
