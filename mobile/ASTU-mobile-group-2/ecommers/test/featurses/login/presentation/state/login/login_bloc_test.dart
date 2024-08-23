



import 'dart:convert';
import 'package:bloc_test/bloc_test.dart';

import 'package:dartz/dartz.dart';
import 'package:ecommers/features/login/data/model/model.dart';
import 'package:ecommers/features/login/domain/entity/login_entity.dart';
import 'package:ecommers/features/login/presentation/state/login/login_bloc.dart';
import 'package:ecommers/features/login/presentation/state/login/login_event.dart';
import 'package:ecommers/features/login/presentation/state/login/login_state.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';

import '../../../../../helper/dummy_data/read_json.dart';
import '../../../../../helper/test_hlper.mocks.dart';


void main() {

  late MockLoginUseCase mockLoginUseCase;
  late LoginBloc loginBloc;

  setUp(() {
    mockLoginUseCase = MockLoginUseCase();
    loginBloc = LoginBloc(loginUseCase: mockLoginUseCase);
  });


  final jsonString = readJson('helper/dummy_data/login_success.json');
  final jsonDecode = json.decode(jsonString);
  final LoginModel loginModel= LoginModel.fromJson(jsonDecode);
  final LoginEntity loginEntity = loginModel.toEntity();

  test(
    'test event of the login event',
    ()  {
      expect(loginBloc.state, isA<LoginIntial>());
    }
    );

  group (
    'test the login bloc',


    () {
      blocTest<LoginBloc, LoginState>(
        'emits [MyState] when MyEvent is added.',
        build: (){
          when(mockLoginUseCase.loginUser('samuel.tolossa@a2sv.org','123qweasdzxc')
          ).thenAnswer((_) async => Right(loginEntity));
          return loginBloc;
          },
          act: (bloc) => bloc.add(LoginRequest(email: 'samuel.tolossa@a2sv.org', password: '123qweasdzxc')),
            expect: () => [
              LoginSuccess(message: 'login Success')
          ]
      );

          
    }
  );
}