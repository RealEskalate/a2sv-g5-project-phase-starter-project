
import '../../domain/entities/log_in.dart';

class LogInModel extends LogInEntity {
  const LogInModel({
    required super.email,
    required super.password,
  });
  Map<String, dynamic> toJson() {
    return {
      'email': email,
      'password': password,
    };
  }
}
