import '../../../../core/exception/exception.dart';
import '../../domain/entities/authenticated_entity.dart';

class AuthenticatedModel extends AuthenticatedEntity {
   AuthenticatedModel({
    required super.token,
  });

  factory AuthenticatedModel.fromJson(Map<String, dynamic> json) {
    try {
      return AuthenticatedModel(
        token: json['access_token'],
      );
    } catch (e) {
      throw JsonParsingException();
    }
  }
}
