import '../../../../core/network/network_info.dart';
import '../../domain/entity/user.dart';
import '../../domain/repository/authentication_repository.dart';
import '../data_source/local_data_source.dart';
import '../data_source/remote_data_source.dart';

class AuthenticationRepositoryImpl extends AuthenticationRepository{
  final NetworkInfo networkInfo;
  final UserLocalDataSource localDataSource;
  final UserRemoteDataSource remoteDataSource;

  AuthenticationRepositoryImpl({required this.networkInfo,required this.localDataSource, required this.remoteDataSource});  

  @override
  Future<bool> isSignedIn() {
    // TODO: implement isSignedIn
    throw UnimplementedError();
  }

  @override

  Future<User> logIn({String? email, String? password}) async {
    // if (await networkInfo.isConnected){
        final user = await remoteDataSource.logIn( email ?? '',  password ?? '');
        localDataSource.cacheUser(user);
        return user;  
    // }else{
    //   throw Exception('No internet connection');
    // }
  }

  @override
  Future<void> logOut() async {
    try {
      await remoteDataSource.logOut();
    }catch (e){
      throw (e);
    }
  }

  @override
  Future<User> signUp({String? email, String? password, String? username})async {
    print("From repo Impl");
    // if (await networkInfo.isConnected){
    try{
      
        final user = await remoteDataSource.signUp(username ?? '', password ?? '', email ?? '');
        localDataSource.cacheUser(user);
        print(user);
        return user;
        }catch(e){
          throw Exception('Server Failure');
        }
  // }else{
  //   throw Exception('No internet connection');
  // }
  
}
}