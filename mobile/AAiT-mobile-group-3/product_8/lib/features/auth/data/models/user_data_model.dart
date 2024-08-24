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
  // add toJson method
  Map<String, dynamic> toJson() {
    return {
      '_id': id,
      'email': email,
      'name': name,
    };
  }

  // to model
  static UserDataModel toModel(UserDataEntity entity) {
    return UserDataModel(
      id: entity.id,
      email: entity.email,
      name: entity.name,
    );
  }
}
