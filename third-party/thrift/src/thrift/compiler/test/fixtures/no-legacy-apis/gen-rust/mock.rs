// @generated by Thrift for thrift/compiler/test/fixtures/no-legacy-apis/src/module.thrift
// This file is probably not the place you want to edit!

//! Mock definitions for `module`.
//!
//! Client mocks. For every service, a struct mock::TheService that implements
//! client::TheService.
//!
//! As an example of the generated API, for the following thrift service:
//!
//! ```thrift
//! service MyService {
//!     FunctionResponse myFunction(
//!         1: FunctionRequest request,
//!     ) throws {
//!         1: StorageException s,
//!         2: NotFoundException n,
//!     ),
//!
//!     // other functions
//! }
//! ```
//!
//! we would end up with this mock object under crate::mock::MyService:
//!
//! ```
//! # const _: &str = stringify! {
//! impl crate::client::MyService for MyService<'mock> {...}
//!
//! pub struct MyService<'mock> {
//!     pub myFunction: myFunction<'mock>,
//!     // ...
//! }
//!
//! impl dyn crate::client::MyService {
//!     pub fn mock<'mock>() -> MyService<'mock>;
//! }
//!
//! impl myFunction<'mock> {
//!     // directly return the given success response
//!     pub fn ret(&self, value: FunctionResponse);
//!
//!     // invoke closure to compute success response
//!     pub fn mock(
//!         &self,
//!         mock: impl FnMut(FunctionRequest) -> FunctionResponse + Send + Sync + 'mock,
//!     );
//!
//!     // invoke closure to compute response
//!     pub fn mock_result(
//!         &self,
//!         mock: impl FnMut(FunctionRequest) -> Result<FunctionResponse, crate::services::MyService::MyFunctionExn> + Send + Sync + 'mock,
//!     );
//!
//!     // return one of the function's declared exceptions
//!     pub fn throw<E>(&self, exception: E)
//!     where
//!         E: Clone + Into<crate::services::MyService::MyFunctionExn> + Send + Sync + 'mock;
//! }
//!
//! impl From<StorageException> for MyFunctionExn {...}
//! impl From<NotFoundException> for MyFunctionExn {...}
//! # };
//! ```
//!
//! The intended usage from a test would be:
//!
//! ```
//! # const _: &str = stringify! {
//! use std::sync::Arc;
//! use thrift_if::client::MyService;
//!
//! #[test]
//! fn test_my_client() {
//!     let mock = Arc::new(<dyn MyService>::mock());
//!
//!     // directly return a success response
//!     let resp = FunctionResponse {...};
//!     mock.myFunction.ret(resp);
//!
//!     // or give a closure to compute the success response
//!     mock.myFunction.mock(|request| FunctionResponse {...});
//!
//!     // or throw one of the function's exceptions
//!     mock.myFunction.throw(StorageException::ItFailed);
//!
//!     // or compute a Result (useful if your exceptions aren't Clone)
//!     mock.myFunction.mock_result(|request| Err(...));
//!
//!     let out = do_the_thing(mock).wait().unwrap();
//!     assert!(out.what_i_expected());
//! }
//!
//! fn do_the_thing(
//!     client: Arc<dyn MyService + Send + Sync + 'static>,
//! ) -> impl Future<Item = Out> {...}
//! # };
//! ```

pub struct MyService<'mock> {
    pub query: r#impl::my_service::query<'mock>,
    _marker: ::std::marker::PhantomData<&'mock ()>,
}

impl dyn super::client::MyService {
    pub fn mock<'mock>() -> MyService<'mock> {
        MyService {
            query: r#impl::my_service::query::unimplemented(),
            _marker: ::std::marker::PhantomData,
        }
    }
}

impl<'mock> super::client::MyService for MyService<'mock> {
    fn query(
        &self,
        arg_u: &crate::types::MyUnion,
    ) -> ::futures::future::BoxFuture<'static, ::std::result::Result<crate::types::MyStruct, crate::errors::my_service::QueryError>> {
        let mut closure = self.query.closure.lock().unwrap();
        let closure: &mut dyn ::std::ops::FnMut(crate::types::MyUnion) -> _ = &mut **closure;
        ::std::boxed::Box::pin(::futures::future::ready(closure(arg_u.clone())))
    }
}

pub mod r#impl {
    pub mod my_service {

        pub struct query<'mock> {
            pub(crate) closure: ::std::sync::Mutex<::std::boxed::Box<
                dyn ::std::ops::FnMut(crate::types::MyUnion) -> ::std::result::Result<
                    crate::types::MyStruct,
                    crate::errors::my_service::QueryError,
                > + ::std::marker::Send + ::std::marker::Sync + 'mock,
            >>,
        }

        #[allow(clippy::redundant_closure)]
        impl<'mock> query<'mock> {
            pub(crate) fn unimplemented() -> Self {
                Self {
                    closure: ::std::sync::Mutex::new(::std::boxed::Box::new(|_: crate::types::MyUnion| panic!(
                        "{}::{} is not mocked",
                        "MyService",
                        "query",
                    ))),
                }
            }

            pub fn ret(&self, value: crate::types::MyStruct) {
                self.mock(move |_: crate::types::MyUnion| value.clone());
            }

            pub fn mock(&self, mut mock: impl ::std::ops::FnMut(crate::types::MyUnion) -> crate::types::MyStruct + ::std::marker::Send + ::std::marker::Sync + 'mock) {
                let mut closure = self.closure.lock().unwrap();
                *closure = ::std::boxed::Box::new(move |u| ::std::result::Result::Ok(mock(u)));
            }

            pub fn mock_result(&self, mut mock: impl ::std::ops::FnMut(crate::types::MyUnion) -> ::std::result::Result<crate::types::MyStruct, crate::errors::my_service::QueryError> + ::std::marker::Send + ::std::marker::Sync + 'mock) {
                let mut closure = self.closure.lock().unwrap();
                *closure = ::std::boxed::Box::new(move |u| mock(u));
            }

            pub fn throw<E>(&self, exception: E)
            where
                E: ::std::convert::Into<crate::errors::my_service::QueryError>,
                E: ::std::clone::Clone + ::std::marker::Send + ::std::marker::Sync + 'mock,
            {
                let mut closure = self.closure.lock().unwrap();
                *closure = ::std::boxed::Box::new(move |_: crate::types::MyUnion| ::std::result::Result::Err(exception.clone().into()));
            }
        }
    }
}
