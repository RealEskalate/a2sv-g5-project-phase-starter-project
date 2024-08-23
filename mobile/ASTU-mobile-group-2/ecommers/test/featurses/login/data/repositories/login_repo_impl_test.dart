

import 'dart:convert';

import 'package:dartz/dartz.dart';
import 'package:ecommers/features/login/data/model/model.dart';
import 'package:ecommers/features/login/data/repositories/login_repo_impl.dart';
import 'package:ecommers/features/login/domain/entity/login_entity.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';

import '../../../../helper/dummy_data/read_json.dart';
import '../../../../helper/test_hlper.mocks.dart';

void main() {

  late MockRemoteDatasourceImpl mockRemoteDatasourceImpl;
  late LoginRepoImpl loginRepoImpl;


  setUp((){
    mockRemoteDatasourceImpl = MockRemoteDatasourceImpl();
    loginRepoImpl = LoginRepoImpl(remoteDatasourceImpl: mockRemoteDatasourceImpl);
  });
  

  final data = readJson('helper/dummy_data/login_success.json');
  final jsonDecode = json.decode(data);
  final LoginModel loginModel = LoginModel.fromJson(jsonDecode);
  final LoginEntity loginEntity = loginModel.toEntity();
  

  group (
    'test the login repo impl',
    () {
      

      test(
      'repo respond must be seccuess',
      () async {
        when(mockRemoteDatasourceImpl.login('samuel.tolossa@a2sv.org', '123qweasdzxc'
        )).thenAnswer((_) async =>Right(loginEntity) );

        // Call the login method and await the result
        final result = await loginRepoImpl.login('samuel.tolossa@a2sv.org', '123qweasdzxc');

        // Check the result type
         expect(result, isA<Right>());
         expect(result.fold((l) => l, (r) => r), isA<LoginEntity>());
      }
      );

    }
  );
}