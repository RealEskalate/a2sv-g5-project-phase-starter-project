import '../../../../core/error/exception.dart';
import '../../domain/entities/user_data_entity.dart';

class UserDataModel extends UserDataEntity{
  UserDataModel({required super.email, required super.name});

  UserDataEntity toEntity() => UserDataEntity(email: email, name: name);

  factory UserDataModel.fromJson(Map<String, dynamic> json){
  try{
    return UserDataModel(email: json['email'], name: json['name']);
  } catch (e){
    throw JsonParsingException();
  }
}
}