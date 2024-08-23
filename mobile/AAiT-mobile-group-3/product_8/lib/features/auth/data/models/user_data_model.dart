import '../../../../core/exception/exception.dart';
import '../../domain/entities/user_data_entity.dart';

class UserDataModel extends UserDataEntity{
   UserDataModel({
    required  email,
    required name,
  }) : super(
    email: email,
    name: name,
  );

 UserDataEntity toEntity() => UserDataEntity(email: email, name: name);

  factory UserDataModel.fromJson(Map<String, dynamic> json) {
    try{return UserDataModel(email: json['email'], name: json['name']);}
    catch (e) {
      throw JsonParsingException();
    }
  }


}

