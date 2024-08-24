// Mocks generated by Mockito 5.4.4 from annotations
// in application1/test/helper/test_helper.dart.
// Do not manually edit this file.

// ignore_for_file: no_leading_underscores_for_library_prefixes
import 'dart:async' as _i11;
import 'dart:convert' as _i35;
import 'dart:typed_data' as _i36;

import 'package:application1/core/error/failure.dart' as _i12;
import 'package:application1/core/network/network_info.dart' as _i16;
import 'package:application1/core/usecase/usecase.dart' as _i20;
import 'package:application1/features/authentication/data/data_sources/local/local_data_source.dart'
    as _i24;
import 'package:application1/features/authentication/data/data_sources/remote/auth_remote_data_source.dart'
    as _i32;
import 'package:application1/features/authentication/data/model/log_in_model.dart'
    as _i34;
import 'package:application1/features/authentication/data/model/sign_up_model.dart'
    as _i33;
import 'package:application1/features/authentication/data/model/user_model.dart'
    as _i9;
import 'package:application1/features/authentication/domain/entities/log_in.dart'
    as _i30;
import 'package:application1/features/authentication/domain/entities/sign_up.dart'
    as _i29;
import 'package:application1/features/authentication/domain/entities/user_data.dart'
    as _i31;
import 'package:application1/features/authentication/domain/repositories/auth_repo.dart'
    as _i28;
import 'package:application1/features/authentication/domain/usecases/get_current_user_usecase.dart'
    as _i5;
import 'package:application1/features/authentication/domain/usecases/log_in_usecase.dart'
    as _i6;
import 'package:application1/features/authentication/domain/usecases/log_out_usecase.dart'
    as _i7;
import 'package:application1/features/authentication/domain/usecases/sign_up_usecase.dart'
    as _i8;
import 'package:application1/features/authentication/presentation/bloc/auth_bloc.dart'
    as _i26;
import 'package:application1/features/product/data/data_sources/local/local_data_source.dart'
    as _i15;
import 'package:application1/features/product/data/data_sources/remote/remote_data_source.dart'
    as _i14;
import 'package:application1/features/product/data/models/product_model.dart'
    as _i3;
import 'package:application1/features/product/domain/entities/product_entity.dart'
    as _i13;
import 'package:application1/features/product/domain/repository/product_repository.dart'
    as _i4;
import 'package:application1/features/product/domain/usecases/add_product_usecase.dart'
    as _i21;
import 'package:application1/features/product/domain/usecases/delete_product_usecase.dart'
    as _i23;
import 'package:application1/features/product/domain/usecases/get_product_usecase.dart'
    as _i18;
import 'package:application1/features/product/domain/usecases/get_products_usecase.dart'
    as _i19;
import 'package:application1/features/product/domain/usecases/update_product_usecase.dart'
    as _i22;
import 'package:bloc/bloc.dart' as _i27;
import 'package:dartz/dartz.dart' as _i2;
import 'package:http/http.dart' as _i10;
import 'package:mockito/mockito.dart' as _i1;
import 'package:mockito/src/dummies.dart' as _i25;
import 'package:shared_preferences/src/shared_preferences_legacy.dart' as _i17;

// ignore_for_file: type=lint
// ignore_for_file: avoid_redundant_argument_values
// ignore_for_file: avoid_setters_without_getters
// ignore_for_file: comment_references
// ignore_for_file: deprecated_member_use
// ignore_for_file: deprecated_member_use_from_same_package
// ignore_for_file: implementation_imports
// ignore_for_file: invalid_use_of_visible_for_testing_member
// ignore_for_file: prefer_const_constructors
// ignore_for_file: unnecessary_parenthesis
// ignore_for_file: camel_case_types
// ignore_for_file: subtype_of_sealed_class

class _FakeEither_0<L, R> extends _i1.SmartFake implements _i2.Either<L, R> {
  _FakeEither_0(
    Object parent,
    Invocation parentInvocation,
  ) : super(
          parent,
          parentInvocation,
        );
}

class _FakeProductModel_1 extends _i1.SmartFake implements _i3.ProductModel {
  _FakeProductModel_1(
    Object parent,
    Invocation parentInvocation,
  ) : super(
          parent,
          parentInvocation,
        );
}

class _FakeProductRepository_2 extends _i1.SmartFake
    implements _i4.ProductRepository {
  _FakeProductRepository_2(
    Object parent,
    Invocation parentInvocation,
  ) : super(
          parent,
          parentInvocation,
        );
}

class _FakeGetCurrentUserUsecase_3 extends _i1.SmartFake
    implements _i5.GetCurrentUserUsecase {
  _FakeGetCurrentUserUsecase_3(
    Object parent,
    Invocation parentInvocation,
  ) : super(
          parent,
          parentInvocation,
        );
}

class _FakeLogInUsecase_4 extends _i1.SmartFake implements _i6.LogInUsecase {
  _FakeLogInUsecase_4(
    Object parent,
    Invocation parentInvocation,
  ) : super(
          parent,
          parentInvocation,
        );
}

class _FakeLogOutUsecase_5 extends _i1.SmartFake implements _i7.LogOutUsecase {
  _FakeLogOutUsecase_5(
    Object parent,
    Invocation parentInvocation,
  ) : super(
          parent,
          parentInvocation,
        );
}

class _FakeSignUpUsecase_6 extends _i1.SmartFake implements _i8.SignUpUsecase {
  _FakeSignUpUsecase_6(
    Object parent,
    Invocation parentInvocation,
  ) : super(
          parent,
          parentInvocation,
        );
}

class _FakeUserModel_7 extends _i1.SmartFake implements _i9.UserModel {
  _FakeUserModel_7(
    Object parent,
    Invocation parentInvocation,
  ) : super(
          parent,
          parentInvocation,
        );
}

class _FakeResponse_8 extends _i1.SmartFake implements _i10.Response {
  _FakeResponse_8(
    Object parent,
    Invocation parentInvocation,
  ) : super(
          parent,
          parentInvocation,
        );
}

class _FakeStreamedResponse_9 extends _i1.SmartFake
    implements _i10.StreamedResponse {
  _FakeStreamedResponse_9(
    Object parent,
    Invocation parentInvocation,
  ) : super(
          parent,
          parentInvocation,
        );
}

/// A class which mocks [ProductRepository].
///
/// See the documentation for Mockito's code generation for more information.
class MockProductRepository extends _i1.Mock implements _i4.ProductRepository {
  MockProductRepository() {
    _i1.throwOnMissingStub(this);
  }

  @override
  _i11.Future<_i2.Either<_i12.Failure, List<_i13.ProductEntity>>>
      getProducts() => (super.noSuchMethod(
            Invocation.method(
              #getProducts,
              [],
            ),
            returnValue: _i11.Future<
                    _i2.Either<_i12.Failure, List<_i13.ProductEntity>>>.value(
                _FakeEither_0<_i12.Failure, List<_i13.ProductEntity>>(
              this,
              Invocation.method(
                #getProducts,
                [],
              ),
            )),
          ) as _i11.Future<_i2.Either<_i12.Failure, List<_i13.ProductEntity>>>);

  @override
  _i11.Future<_i2.Either<_i12.Failure, _i13.ProductEntity>> getProduct(
          String? id) =>
      (super.noSuchMethod(
        Invocation.method(
          #getProduct,
          [id],
        ),
        returnValue:
            _i11.Future<_i2.Either<_i12.Failure, _i13.ProductEntity>>.value(
                _FakeEither_0<_i12.Failure, _i13.ProductEntity>(
          this,
          Invocation.method(
            #getProduct,
            [id],
          ),
        )),
      ) as _i11.Future<_i2.Either<_i12.Failure, _i13.ProductEntity>>);

  @override
  _i11.Future<_i2.Either<_i12.Failure, _i13.ProductEntity>> updateProduct(
          _i13.ProductEntity? product) =>
      (super.noSuchMethod(
        Invocation.method(
          #updateProduct,
          [product],
        ),
        returnValue:
            _i11.Future<_i2.Either<_i12.Failure, _i13.ProductEntity>>.value(
                _FakeEither_0<_i12.Failure, _i13.ProductEntity>(
          this,
          Invocation.method(
            #updateProduct,
            [product],
          ),
        )),
      ) as _i11.Future<_i2.Either<_i12.Failure, _i13.ProductEntity>>);

  @override
  _i11.Future<_i2.Either<_i12.Failure, bool>> deleteProduct(String? id) =>
      (super.noSuchMethod(
        Invocation.method(
          #deleteProduct,
          [id],
        ),
        returnValue: _i11.Future<_i2.Either<_i12.Failure, bool>>.value(
            _FakeEither_0<_i12.Failure, bool>(
          this,
          Invocation.method(
            #deleteProduct,
            [id],
          ),
        )),
      ) as _i11.Future<_i2.Either<_i12.Failure, bool>>);

  @override
  _i11.Future<_i2.Either<_i12.Failure, _i13.ProductEntity>> addProduct(
          _i13.ProductEntity? product) =>
      (super.noSuchMethod(
        Invocation.method(
          #addProduct,
          [product],
        ),
        returnValue:
            _i11.Future<_i2.Either<_i12.Failure, _i13.ProductEntity>>.value(
                _FakeEither_0<_i12.Failure, _i13.ProductEntity>(
          this,
          Invocation.method(
            #addProduct,
            [product],
          ),
        )),
      ) as _i11.Future<_i2.Either<_i12.Failure, _i13.ProductEntity>>);
}

/// A class which mocks [ProductRemoteDataSource].
///
/// See the documentation for Mockito's code generation for more information.
class MockProductRemoteDataSource extends _i1.Mock
    implements _i14.ProductRemoteDataSource {
  MockProductRemoteDataSource() {
    _i1.throwOnMissingStub(this);
  }

  @override
  _i11.Future<_i3.ProductModel> getProduct(String? id) => (super.noSuchMethod(
        Invocation.method(
          #getProduct,
          [id],
        ),
        returnValue: _i11.Future<_i3.ProductModel>.value(_FakeProductModel_1(
          this,
          Invocation.method(
            #getProduct,
            [id],
          ),
        )),
      ) as _i11.Future<_i3.ProductModel>);

  @override
  _i11.Future<List<_i3.ProductModel>> getProducts() => (super.noSuchMethod(
        Invocation.method(
          #getProducts,
          [],
        ),
        returnValue:
            _i11.Future<List<_i3.ProductModel>>.value(<_i3.ProductModel>[]),
      ) as _i11.Future<List<_i3.ProductModel>>);

  @override
  _i11.Future<bool> deleteProduct(String? id) => (super.noSuchMethod(
        Invocation.method(
          #deleteProduct,
          [id],
        ),
        returnValue: _i11.Future<bool>.value(false),
      ) as _i11.Future<bool>);

  @override
  _i11.Future<_i3.ProductModel> updateProduct(_i3.ProductModel? product) =>
      (super.noSuchMethod(
        Invocation.method(
          #updateProduct,
          [product],
        ),
        returnValue: _i11.Future<_i3.ProductModel>.value(_FakeProductModel_1(
          this,
          Invocation.method(
            #updateProduct,
            [product],
          ),
        )),
      ) as _i11.Future<_i3.ProductModel>);

  @override
  _i11.Future<_i3.ProductModel> addProduct(_i3.ProductModel? product) =>
      (super.noSuchMethod(
        Invocation.method(
          #addProduct,
          [product],
        ),
        returnValue: _i11.Future<_i3.ProductModel>.value(_FakeProductModel_1(
          this,
          Invocation.method(
            #addProduct,
            [product],
          ),
        )),
      ) as _i11.Future<_i3.ProductModel>);
}

/// A class which mocks [ProductLocalDataSource].
///
/// See the documentation for Mockito's code generation for more information.
class MockProductLocalDataSource extends _i1.Mock
    implements _i15.ProductLocalDataSource {
  MockProductLocalDataSource() {
    _i1.throwOnMissingStub(this);
  }

  @override
  _i11.Future<bool> cacheProducts(List<_i3.ProductModel>? products) =>
      (super.noSuchMethod(
        Invocation.method(
          #cacheProducts,
          [products],
        ),
        returnValue: _i11.Future<bool>.value(false),
      ) as _i11.Future<bool>);

  @override
  _i11.Future<List<_i3.ProductModel>> getProducts() => (super.noSuchMethod(
        Invocation.method(
          #getProducts,
          [],
        ),
        returnValue:
            _i11.Future<List<_i3.ProductModel>>.value(<_i3.ProductModel>[]),
      ) as _i11.Future<List<_i3.ProductModel>>);
}

/// A class which mocks [NetworkInfo].
///
/// See the documentation for Mockito's code generation for more information.
class MockNetworkInfo extends _i1.Mock implements _i16.NetworkInfo {
  MockNetworkInfo() {
    _i1.throwOnMissingStub(this);
  }

  @override
  _i11.Future<bool> get isConnected => (super.noSuchMethod(
        Invocation.getter(#isConnected),
        returnValue: _i11.Future<bool>.value(false),
      ) as _i11.Future<bool>);
}

/// A class which mocks [SharedPreferences].
///
/// See the documentation for Mockito's code generation for more information.
class MockSharedPreferences extends _i1.Mock implements _i17.SharedPreferences {
  MockSharedPreferences() {
    _i1.throwOnMissingStub(this);
  }

  @override
  Set<String> getKeys() => (super.noSuchMethod(
        Invocation.method(
          #getKeys,
          [],
        ),
        returnValue: <String>{},
      ) as Set<String>);

  @override
  Object? get(String? key) => (super.noSuchMethod(Invocation.method(
        #get,
        [key],
      )) as Object?);

  @override
  bool? getBool(String? key) => (super.noSuchMethod(Invocation.method(
        #getBool,
        [key],
      )) as bool?);

  @override
  int? getInt(String? key) => (super.noSuchMethod(Invocation.method(
        #getInt,
        [key],
      )) as int?);

  @override
  double? getDouble(String? key) => (super.noSuchMethod(Invocation.method(
        #getDouble,
        [key],
      )) as double?);

  @override
  String? getString(String? key) => (super.noSuchMethod(Invocation.method(
        #getString,
        [key],
      )) as String?);

  @override
  bool containsKey(String? key) => (super.noSuchMethod(
        Invocation.method(
          #containsKey,
          [key],
        ),
        returnValue: false,
      ) as bool);

  @override
  List<String>? getStringList(String? key) =>
      (super.noSuchMethod(Invocation.method(
        #getStringList,
        [key],
      )) as List<String>?);

  @override
  _i11.Future<bool> setBool(
    String? key,
    bool? value,
  ) =>
      (super.noSuchMethod(
        Invocation.method(
          #setBool,
          [
            key,
            value,
          ],
        ),
        returnValue: _i11.Future<bool>.value(false),
      ) as _i11.Future<bool>);

  @override
  _i11.Future<bool> setInt(
    String? key,
    int? value,
  ) =>
      (super.noSuchMethod(
        Invocation.method(
          #setInt,
          [
            key,
            value,
          ],
        ),
        returnValue: _i11.Future<bool>.value(false),
      ) as _i11.Future<bool>);

  @override
  _i11.Future<bool> setDouble(
    String? key,
    double? value,
  ) =>
      (super.noSuchMethod(
        Invocation.method(
          #setDouble,
          [
            key,
            value,
          ],
        ),
        returnValue: _i11.Future<bool>.value(false),
      ) as _i11.Future<bool>);

  @override
  _i11.Future<bool> setString(
    String? key,
    String? value,
  ) =>
      (super.noSuchMethod(
        Invocation.method(
          #setString,
          [
            key,
            value,
          ],
        ),
        returnValue: _i11.Future<bool>.value(false),
      ) as _i11.Future<bool>);

  @override
  _i11.Future<bool> setStringList(
    String? key,
    List<String>? value,
  ) =>
      (super.noSuchMethod(
        Invocation.method(
          #setStringList,
          [
            key,
            value,
          ],
        ),
        returnValue: _i11.Future<bool>.value(false),
      ) as _i11.Future<bool>);

  @override
  _i11.Future<bool> remove(String? key) => (super.noSuchMethod(
        Invocation.method(
          #remove,
          [key],
        ),
        returnValue: _i11.Future<bool>.value(false),
      ) as _i11.Future<bool>);

  @override
  _i11.Future<bool> commit() => (super.noSuchMethod(
        Invocation.method(
          #commit,
          [],
        ),
        returnValue: _i11.Future<bool>.value(false),
      ) as _i11.Future<bool>);

  @override
  _i11.Future<bool> clear() => (super.noSuchMethod(
        Invocation.method(
          #clear,
          [],
        ),
        returnValue: _i11.Future<bool>.value(false),
      ) as _i11.Future<bool>);

  @override
  _i11.Future<void> reload() => (super.noSuchMethod(
        Invocation.method(
          #reload,
          [],
        ),
        returnValue: _i11.Future<void>.value(),
        returnValueForMissingStub: _i11.Future<void>.value(),
      ) as _i11.Future<void>);
}

/// A class which mocks [GetProductUsecase].
///
/// See the documentation for Mockito's code generation for more information.
class MockGetProductUsecase extends _i1.Mock implements _i18.GetProductUsecase {
  MockGetProductUsecase() {
    _i1.throwOnMissingStub(this);
  }

  @override
  _i4.ProductRepository get repository => (super.noSuchMethod(
        Invocation.getter(#repository),
        returnValue: _FakeProductRepository_2(
          this,
          Invocation.getter(#repository),
        ),
      ) as _i4.ProductRepository);

  @override
  _i11.Future<_i2.Either<_i12.Failure, _i13.ProductEntity>> call(
          _i18.GetParams? params) =>
      (super.noSuchMethod(
        Invocation.method(
          #call,
          [params],
        ),
        returnValue:
            _i11.Future<_i2.Either<_i12.Failure, _i13.ProductEntity>>.value(
                _FakeEither_0<_i12.Failure, _i13.ProductEntity>(
          this,
          Invocation.method(
            #call,
            [params],
          ),
        )),
      ) as _i11.Future<_i2.Either<_i12.Failure, _i13.ProductEntity>>);
}

/// A class which mocks [GetProductsUsecase].
///
/// See the documentation for Mockito's code generation for more information.
class MockGetProductsUsecase extends _i1.Mock
    implements _i19.GetProductsUsecase {
  MockGetProductsUsecase() {
    _i1.throwOnMissingStub(this);
  }

  @override
  _i4.ProductRepository get repository => (super.noSuchMethod(
        Invocation.getter(#repository),
        returnValue: _FakeProductRepository_2(
          this,
          Invocation.getter(#repository),
        ),
      ) as _i4.ProductRepository);

  @override
  _i11.Future<_i2.Either<_i12.Failure, List<_i13.ProductEntity>>> call(
          _i20.NoParams? params) =>
      (super.noSuchMethod(
        Invocation.method(
          #call,
          [params],
        ),
        returnValue: _i11
            .Future<_i2.Either<_i12.Failure, List<_i13.ProductEntity>>>.value(
            _FakeEither_0<_i12.Failure, List<_i13.ProductEntity>>(
          this,
          Invocation.method(
            #call,
            [params],
          ),
        )),
      ) as _i11.Future<_i2.Either<_i12.Failure, List<_i13.ProductEntity>>>);
}

/// A class which mocks [AddProductUsecase].
///
/// See the documentation for Mockito's code generation for more information.
class MockAddProductUsecase extends _i1.Mock implements _i21.AddProductUsecase {
  MockAddProductUsecase() {
    _i1.throwOnMissingStub(this);
  }

  @override
  _i4.ProductRepository get productRepository => (super.noSuchMethod(
        Invocation.getter(#productRepository),
        returnValue: _FakeProductRepository_2(
          this,
          Invocation.getter(#productRepository),
        ),
      ) as _i4.ProductRepository);

  @override
  _i11.Future<_i2.Either<_i12.Failure, _i13.ProductEntity>> call(
          _i21.CreateParams? params) =>
      (super.noSuchMethod(
        Invocation.method(
          #call,
          [params],
        ),
        returnValue:
            _i11.Future<_i2.Either<_i12.Failure, _i13.ProductEntity>>.value(
                _FakeEither_0<_i12.Failure, _i13.ProductEntity>(
          this,
          Invocation.method(
            #call,
            [params],
          ),
        )),
      ) as _i11.Future<_i2.Either<_i12.Failure, _i13.ProductEntity>>);
}

/// A class which mocks [UpdateProductUsecase].
///
/// See the documentation for Mockito's code generation for more information.
class MockUpdateProductUsecase extends _i1.Mock
    implements _i22.UpdateProductUsecase {
  MockUpdateProductUsecase() {
    _i1.throwOnMissingStub(this);
  }

  @override
  _i4.ProductRepository get repository => (super.noSuchMethod(
        Invocation.getter(#repository),
        returnValue: _FakeProductRepository_2(
          this,
          Invocation.getter(#repository),
        ),
      ) as _i4.ProductRepository);

  @override
  _i11.Future<_i2.Either<_i12.Failure, _i13.ProductEntity>> call(
          _i22.UpdateParams? params) =>
      (super.noSuchMethod(
        Invocation.method(
          #call,
          [params],
        ),
        returnValue:
            _i11.Future<_i2.Either<_i12.Failure, _i13.ProductEntity>>.value(
                _FakeEither_0<_i12.Failure, _i13.ProductEntity>(
          this,
          Invocation.method(
            #call,
            [params],
          ),
        )),
      ) as _i11.Future<_i2.Either<_i12.Failure, _i13.ProductEntity>>);
}

/// A class which mocks [DeleteProductUsecase].
///
/// See the documentation for Mockito's code generation for more information.
class MockDeleteProductUsecase extends _i1.Mock
    implements _i23.DeleteProductUsecase {
  MockDeleteProductUsecase() {
    _i1.throwOnMissingStub(this);
  }

  @override
  _i4.ProductRepository get repository => (super.noSuchMethod(
        Invocation.getter(#repository),
        returnValue: _FakeProductRepository_2(
          this,
          Invocation.getter(#repository),
        ),
      ) as _i4.ProductRepository);

  @override
  _i11.Future<_i2.Either<_i12.Failure, bool>> call(_i23.DeleteParams? params) =>
      (super.noSuchMethod(
        Invocation.method(
          #call,
          [params],
        ),
        returnValue: _i11.Future<_i2.Either<_i12.Failure, bool>>.value(
            _FakeEither_0<_i12.Failure, bool>(
          this,
          Invocation.method(
            #call,
            [params],
          ),
        )),
      ) as _i11.Future<_i2.Either<_i12.Failure, bool>>);
}

/// A class which mocks [AuthLocalDataSource].
///
/// See the documentation for Mockito's code generation for more information.
class MockAuthLocalDataSource extends _i1.Mock
    implements _i24.AuthLocalDataSource {
  MockAuthLocalDataSource() {
    _i1.throwOnMissingStub(this);
  }

  @override
  _i11.Future<bool> cacheToken(String? token) => (super.noSuchMethod(
        Invocation.method(
          #cacheToken,
          [token],
        ),
        returnValue: _i11.Future<bool>.value(false),
      ) as _i11.Future<bool>);

  @override
  _i11.Future<String> getToken() => (super.noSuchMethod(
        Invocation.method(
          #getToken,
          [],
        ),
        returnValue: _i11.Future<String>.value(_i25.dummyValue<String>(
          this,
          Invocation.method(
            #getToken,
            [],
          ),
        )),
      ) as _i11.Future<String>);

  @override
  _i11.Future<void> removeToken() => (super.noSuchMethod(
        Invocation.method(
          #removeToken,
          [],
        ),
        returnValue: _i11.Future<void>.value(),
        returnValueForMissingStub: _i11.Future<void>.value(),
      ) as _i11.Future<void>);
}

/// A class which mocks [AuthBloc].
///
/// See the documentation for Mockito's code generation for more information.
class MockAuthBloc extends _i1.Mock implements _i26.AuthBloc {
  MockAuthBloc() {
    _i1.throwOnMissingStub(this);
  }

  @override
  _i5.GetCurrentUserUsecase get getCurrentUserUsecase => (super.noSuchMethod(
        Invocation.getter(#getCurrentUserUsecase),
        returnValue: _FakeGetCurrentUserUsecase_3(
          this,
          Invocation.getter(#getCurrentUserUsecase),
        ),
      ) as _i5.GetCurrentUserUsecase);

  @override
  _i6.LogInUsecase get logInUsecase => (super.noSuchMethod(
        Invocation.getter(#logInUsecase),
        returnValue: _FakeLogInUsecase_4(
          this,
          Invocation.getter(#logInUsecase),
        ),
      ) as _i6.LogInUsecase);

  @override
  _i7.LogOutUsecase get logOutUsecase => (super.noSuchMethod(
        Invocation.getter(#logOutUsecase),
        returnValue: _FakeLogOutUsecase_5(
          this,
          Invocation.getter(#logOutUsecase),
        ),
      ) as _i7.LogOutUsecase);

  @override
  _i8.SignUpUsecase get signUpUsecase => (super.noSuchMethod(
        Invocation.getter(#signUpUsecase),
        returnValue: _FakeSignUpUsecase_6(
          this,
          Invocation.getter(#signUpUsecase),
        ),
      ) as _i8.SignUpUsecase);

  @override
  _i26.AuthState get state => (super.noSuchMethod(
        Invocation.getter(#state),
        returnValue: _i25.dummyValue<_i26.AuthState>(
          this,
          Invocation.getter(#state),
        ),
      ) as _i26.AuthState);

  @override
  _i11.Stream<_i26.AuthState> get stream => (super.noSuchMethod(
        Invocation.getter(#stream),
        returnValue: _i11.Stream<_i26.AuthState>.empty(),
      ) as _i11.Stream<_i26.AuthState>);

  @override
  bool get isClosed => (super.noSuchMethod(
        Invocation.getter(#isClosed),
        returnValue: false,
      ) as bool);

  @override
  void add(_i26.AuthEvent? event) => super.noSuchMethod(
        Invocation.method(
          #add,
          [event],
        ),
        returnValueForMissingStub: null,
      );

  @override
  void onEvent(_i26.AuthEvent? event) => super.noSuchMethod(
        Invocation.method(
          #onEvent,
          [event],
        ),
        returnValueForMissingStub: null,
      );

  @override
  void emit(_i26.AuthState? state) => super.noSuchMethod(
        Invocation.method(
          #emit,
          [state],
        ),
        returnValueForMissingStub: null,
      );

  @override
  void on<E extends _i26.AuthEvent>(
    _i27.EventHandler<E, _i26.AuthState>? handler, {
    _i27.EventTransformer<E>? transformer,
  }) =>
      super.noSuchMethod(
        Invocation.method(
          #on,
          [handler],
          {#transformer: transformer},
        ),
        returnValueForMissingStub: null,
      );

  @override
  void onTransition(
          _i27.Transition<_i26.AuthEvent, _i26.AuthState>? transition) =>
      super.noSuchMethod(
        Invocation.method(
          #onTransition,
          [transition],
        ),
        returnValueForMissingStub: null,
      );

  @override
  _i11.Future<void> close() => (super.noSuchMethod(
        Invocation.method(
          #close,
          [],
        ),
        returnValue: _i11.Future<void>.value(),
        returnValueForMissingStub: _i11.Future<void>.value(),
      ) as _i11.Future<void>);

  @override
  void onChange(_i27.Change<_i26.AuthState>? change) => super.noSuchMethod(
        Invocation.method(
          #onChange,
          [change],
        ),
        returnValueForMissingStub: null,
      );

  @override
  void addError(
    Object? error, [
    StackTrace? stackTrace,
  ]) =>
      super.noSuchMethod(
        Invocation.method(
          #addError,
          [
            error,
            stackTrace,
          ],
        ),
        returnValueForMissingStub: null,
      );

  @override
  void onError(
    Object? error,
    StackTrace? stackTrace,
  ) =>
      super.noSuchMethod(
        Invocation.method(
          #onError,
          [
            error,
            stackTrace,
          ],
        ),
        returnValueForMissingStub: null,
      );
}

/// A class which mocks [AuthRepository].
///
/// See the documentation for Mockito's code generation for more information.
class MockAuthRepository extends _i1.Mock implements _i28.AuthRepository {
  MockAuthRepository() {
    _i1.throwOnMissingStub(this);
  }

  @override
  _i11.Future<_i2.Either<_i12.Failure, void>> signUp(
          _i29.SignUpEntity? signUpEntity) =>
      (super.noSuchMethod(
        Invocation.method(
          #signUp,
          [signUpEntity],
        ),
        returnValue: _i11.Future<_i2.Either<_i12.Failure, void>>.value(
            _FakeEither_0<_i12.Failure, void>(
          this,
          Invocation.method(
            #signUp,
            [signUpEntity],
          ),
        )),
      ) as _i11.Future<_i2.Either<_i12.Failure, void>>);

  @override
  _i11.Future<_i2.Either<_i12.Failure, void>> logIn(
          _i30.LogInEntity? logInEntity) =>
      (super.noSuchMethod(
        Invocation.method(
          #logIn,
          [logInEntity],
        ),
        returnValue: _i11.Future<_i2.Either<_i12.Failure, void>>.value(
            _FakeEither_0<_i12.Failure, void>(
          this,
          Invocation.method(
            #logIn,
            [logInEntity],
          ),
        )),
      ) as _i11.Future<_i2.Either<_i12.Failure, void>>);

  @override
  _i11.Future<_i2.Either<_i12.Failure, void>> logOut() => (super.noSuchMethod(
        Invocation.method(
          #logOut,
          [],
        ),
        returnValue: _i11.Future<_i2.Either<_i12.Failure, void>>.value(
            _FakeEither_0<_i12.Failure, void>(
          this,
          Invocation.method(
            #logOut,
            [],
          ),
        )),
      ) as _i11.Future<_i2.Either<_i12.Failure, void>>);

  @override
  _i11.Future<_i2.Either<_i12.Failure, _i31.UserEntity>> getCurrentUser() =>
      (super.noSuchMethod(
        Invocation.method(
          #getCurrentUser,
          [],
        ),
        returnValue:
            _i11.Future<_i2.Either<_i12.Failure, _i31.UserEntity>>.value(
                _FakeEither_0<_i12.Failure, _i31.UserEntity>(
          this,
          Invocation.method(
            #getCurrentUser,
            [],
          ),
        )),
      ) as _i11.Future<_i2.Either<_i12.Failure, _i31.UserEntity>>);
}

/// A class which mocks [AuthRemoteDataSource].
///
/// See the documentation for Mockito's code generation for more information.
class MockAuthRemoteDataSource extends _i1.Mock
    implements _i32.AuthRemoteDataSource {
  MockAuthRemoteDataSource() {
    _i1.throwOnMissingStub(this);
  }

  @override
  _i11.Future<void> signUp(_i33.SignUpModel? signUpModel) =>
      (super.noSuchMethod(
        Invocation.method(
          #signUp,
          [signUpModel],
        ),
        returnValue: _i11.Future<void>.value(),
        returnValueForMissingStub: _i11.Future<void>.value(),
      ) as _i11.Future<void>);

  @override
  _i11.Future<void> logIn(_i34.LogInModel? logInModel) => (super.noSuchMethod(
        Invocation.method(
          #logIn,
          [logInModel],
        ),
        returnValue: _i11.Future<void>.value(),
        returnValueForMissingStub: _i11.Future<void>.value(),
      ) as _i11.Future<void>);

  @override
  _i11.Future<void> logOut() => (super.noSuchMethod(
        Invocation.method(
          #logOut,
          [],
        ),
        returnValue: _i11.Future<void>.value(),
        returnValueForMissingStub: _i11.Future<void>.value(),
      ) as _i11.Future<void>);

  @override
  _i11.Future<_i9.UserModel> getCurrentUser() => (super.noSuchMethod(
        Invocation.method(
          #getCurrentUser,
          [],
        ),
        returnValue: _i11.Future<_i9.UserModel>.value(_FakeUserModel_7(
          this,
          Invocation.method(
            #getCurrentUser,
            [],
          ),
        )),
      ) as _i11.Future<_i9.UserModel>);
}

/// A class which mocks [Client].
///
/// See the documentation for Mockito's code generation for more information.
class MockHttpClient extends _i1.Mock implements _i10.Client {
  MockHttpClient() {
    _i1.throwOnMissingStub(this);
  }

  @override
  _i11.Future<_i10.Response> head(
    Uri? url, {
    Map<String, String>? headers,
  }) =>
      (super.noSuchMethod(
        Invocation.method(
          #head,
          [url],
          {#headers: headers},
        ),
        returnValue: _i11.Future<_i10.Response>.value(_FakeResponse_8(
          this,
          Invocation.method(
            #head,
            [url],
            {#headers: headers},
          ),
        )),
      ) as _i11.Future<_i10.Response>);

  @override
  _i11.Future<_i10.Response> get(
    Uri? url, {
    Map<String, String>? headers,
  }) =>
      (super.noSuchMethod(
        Invocation.method(
          #get,
          [url],
          {#headers: headers},
        ),
        returnValue: _i11.Future<_i10.Response>.value(_FakeResponse_8(
          this,
          Invocation.method(
            #get,
            [url],
            {#headers: headers},
          ),
        )),
      ) as _i11.Future<_i10.Response>);

  @override
  _i11.Future<_i10.Response> post(
    Uri? url, {
    Map<String, String>? headers,
    Object? body,
    _i35.Encoding? encoding,
  }) =>
      (super.noSuchMethod(
        Invocation.method(
          #post,
          [url],
          {
            #headers: headers,
            #body: body,
            #encoding: encoding,
          },
        ),
        returnValue: _i11.Future<_i10.Response>.value(_FakeResponse_8(
          this,
          Invocation.method(
            #post,
            [url],
            {
              #headers: headers,
              #body: body,
              #encoding: encoding,
            },
          ),
        )),
      ) as _i11.Future<_i10.Response>);

  @override
  _i11.Future<_i10.Response> put(
    Uri? url, {
    Map<String, String>? headers,
    Object? body,
    _i35.Encoding? encoding,
  }) =>
      (super.noSuchMethod(
        Invocation.method(
          #put,
          [url],
          {
            #headers: headers,
            #body: body,
            #encoding: encoding,
          },
        ),
        returnValue: _i11.Future<_i10.Response>.value(_FakeResponse_8(
          this,
          Invocation.method(
            #put,
            [url],
            {
              #headers: headers,
              #body: body,
              #encoding: encoding,
            },
          ),
        )),
      ) as _i11.Future<_i10.Response>);

  @override
  _i11.Future<_i10.Response> patch(
    Uri? url, {
    Map<String, String>? headers,
    Object? body,
    _i35.Encoding? encoding,
  }) =>
      (super.noSuchMethod(
        Invocation.method(
          #patch,
          [url],
          {
            #headers: headers,
            #body: body,
            #encoding: encoding,
          },
        ),
        returnValue: _i11.Future<_i10.Response>.value(_FakeResponse_8(
          this,
          Invocation.method(
            #patch,
            [url],
            {
              #headers: headers,
              #body: body,
              #encoding: encoding,
            },
          ),
        )),
      ) as _i11.Future<_i10.Response>);

  @override
  _i11.Future<_i10.Response> delete(
    Uri? url, {
    Map<String, String>? headers,
    Object? body,
    _i35.Encoding? encoding,
  }) =>
      (super.noSuchMethod(
        Invocation.method(
          #delete,
          [url],
          {
            #headers: headers,
            #body: body,
            #encoding: encoding,
          },
        ),
        returnValue: _i11.Future<_i10.Response>.value(_FakeResponse_8(
          this,
          Invocation.method(
            #delete,
            [url],
            {
              #headers: headers,
              #body: body,
              #encoding: encoding,
            },
          ),
        )),
      ) as _i11.Future<_i10.Response>);

  @override
  _i11.Future<String> read(
    Uri? url, {
    Map<String, String>? headers,
  }) =>
      (super.noSuchMethod(
        Invocation.method(
          #read,
          [url],
          {#headers: headers},
        ),
        returnValue: _i11.Future<String>.value(_i25.dummyValue<String>(
          this,
          Invocation.method(
            #read,
            [url],
            {#headers: headers},
          ),
        )),
      ) as _i11.Future<String>);

  @override
  _i11.Future<_i36.Uint8List> readBytes(
    Uri? url, {
    Map<String, String>? headers,
  }) =>
      (super.noSuchMethod(
        Invocation.method(
          #readBytes,
          [url],
          {#headers: headers},
        ),
        returnValue: _i11.Future<_i36.Uint8List>.value(_i36.Uint8List(0)),
      ) as _i11.Future<_i36.Uint8List>);

  @override
  _i11.Future<_i10.StreamedResponse> send(_i10.BaseRequest? request) =>
      (super.noSuchMethod(
        Invocation.method(
          #send,
          [request],
        ),
        returnValue:
            _i11.Future<_i10.StreamedResponse>.value(_FakeStreamedResponse_9(
          this,
          Invocation.method(
            #send,
            [request],
          ),
        )),
      ) as _i11.Future<_i10.StreamedResponse>);

  @override
  void close() => super.noSuchMethod(
        Invocation.method(
          #close,
          [],
        ),
        returnValueForMissingStub: null,
      );
}