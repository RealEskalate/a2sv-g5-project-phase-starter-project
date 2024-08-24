import '../../../../core/exception/exception.dart';
import '../../domain/entities/user_data_entity.dart';

class UserDataModel extends UserDataEntity {
  const UserDataModel({
    required super.id,
    required super.email,
    required super.name,
  });

  UserDataEntity toEntity() => UserDataEntity(
        email: email,
        name: name,
        id: id,
      );

  factory UserDataModel.fromJson(Map<String, dynamic> json) {
    try {
      return UserDataModel(
        id: json['_id'],
        email: json['email'],
        name: json['name'],
      );
    } catch (e) {
      throw JsonParsingException();
    }
  }
}
